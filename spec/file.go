package spec

import (
	"bytes"
	"fmt"
	"strings"

	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

// File spec wraps FileDescriptorProto
// and this spec will be passed in all other specs in order to get
// filename, package name, etc...
type File struct {
	descriptor *descriptorpb.FileDescriptorProto
	comments   Comments

	messages []*Message
	services []*Service
	enums    []*Enum

	isCamel bool

	CompilerVersion *pluginpb.Version
}

func NewFile(
	d *descriptorpb.FileDescriptorProto,
	cv *pluginpb.Version,
	isCamel bool,
) *File {

	f := &File{
		CompilerVersion: cv,
		descriptor:      d,
		comments:        makeComments(d),

		services: make([]*Service, 0),
		messages: make([]*Message, 0),
		enums:    make([]*Enum, 0),
		isCamel:  isCamel,
	}
	for i, s := range d.GetService() {
		f.services = append(f.services, NewService(s, f, 6, i)) // nolint: gomnd
	}
	for i, m := range d.GetMessageType() {
		f.messages = append(f.messages, f.messagesRecursive(m, []string{}, 4, i)...) // nolint: gomnd
	}
	for i, e := range d.GetEnumType() {
		f.enums = append(f.enums, NewEnum(e, f, []string{}, 5, i)) // nolint: gomnd
	}
	return f
}

func (f *File) Services() []*Service {
	return f.services
}

func (f *File) Messages() []*Message {
	return f.messages
}

func (f *File) Enums() []*Enum {
	return f.enums
}

func (f *File) messagesRecursive(d *descriptorpb.DescriptorProto, prefix []string, paths ...int) []*Message {
	m := NewMessage(d, f, prefix, f.isCamel, paths...)

	// If message is map_entry, assign all fields as "required"
	if opt := d.GetOptions(); opt != nil && opt.GetMapEntry() {
		m.setRequiredFields()
	}
	messages := []*Message{m}

	prefix = append(prefix, d.GetName())

	// Include enums defined within message
	for i, e := range d.GetEnumType() {
		p := make([]int, len(paths))
		copy(p, paths)
		f.enums = append(f.enums, NewEnum(e, f, prefix, append(p, 5, i)...)) // nolint: gomnd
	}

	for i, m := range d.GetNestedType() {
		p := make([]int, len(paths))
		copy(p, paths)
		messages = append(
			messages,
			f.messagesRecursive(m, prefix, append(p, 3, i)...)..., // nolint: gomnd
		)
	}
	return messages
}

// Package is a package name retrieved from protobuf's package keyword
func (f *File) Package() string {
	return f.descriptor.GetPackage()
}

// GoPackage will search for an option named go_package and, if found, returns it.
// Otherwise, it calls Package
func (f *File) GoPackage() string {
	var pkgName string
	if opt := f.descriptor.GetOptions(); opt == nil {
		pkgName = f.Package()
	} else if p := opt.GetGoPackage(); p == "" {
		pkgName = f.Package()
	} else {
		pkgName = p
	}
	return pkgName
}

func (f *File) Filename() string {
	return f.descriptor.GetName()
}

func (f *File) getComment(paths []int) string {
	b := new(bytes.Buffer)
	for _, p := range paths {
		b.WriteString(fmt.Sprint(p))
	}

	if v, ok := f.comments[b.String()]; ok {
		return strings.ReplaceAll(strings.TrimSpace(v), "`", "`+\"`\"+`")
	}
	return ""
}
