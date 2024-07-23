package spec

import (
	graphqlv1 "github.com/nebucloud/nebucloud-gateway/gen/go/graphql/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Service spec wraps ServiceDescriptorProto with GraphqlService option.
type Service struct {
	descriptor *descriptorpb.ServiceDescriptorProto
	Option     *graphqlv1.GraphqlService
	*File
	paths   []int
	methods []*Method

	Queries   []*Query
	Mutations []*Mutation
}

func NewService(
	d *descriptorpb.ServiceDescriptorProto,
	f *File,
	paths ...int,
) *Service {

	var o *graphqlv1.GraphqlService
	if opts := d.GetOptions(); opts != nil {
		ext := proto.GetExtension(opts, graphqlv1.E_Service)
		if service, ok := ext.(*graphqlv1.GraphqlService); ok {
			o = service
		}
	}

	s := &Service{
		descriptor: d,
		Option:     o,
		File:       f,
		paths:      paths,
		methods:    make([]*Method, 0),
		Queries:    make([]*Query, 0),
		Mutations:  make([]*Mutation, 0),
	}

	for i, m := range d.GetMethod() {
		ps := make([]int, len(paths))
		copy(ps, paths)
		s.methods = append(s.methods, NewMethod(m, s, append(ps, 4, i)...)) // nolint: gomnd
	}
	return s
}

func (s *Service) Comment() string {
	return s.File.getComment(s.paths)
}

func (s *Service) Name() string {
	return s.descriptor.GetName()
}

func (s *Service) Methods() []*Method {
	return s.methods
}

func (s *Service) Host() string {
	if s.Option == nil {
		return ""
	}
	return s.Option.GetHost()
}

func (s *Service) Insecure() bool {
	if s.Option == nil {
		return false
	}
	return s.Option.GetInsecure()
}
