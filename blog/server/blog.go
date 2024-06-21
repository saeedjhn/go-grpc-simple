package main

import pb "github.com/saeedjhn/go-grpc/blog/goproto"

var BlogList = make([]*Blog, 0)

type Blog struct {
	ID       string
	AuthorID string
	Title    string
	Content  string
}

func mapGetBlogResponseToProtobuf(blog *Blog) *pb.Blog {
	return &pb.Blog{
		Id:       blog.ID,
		AuthorId: blog.AuthorID,
		Title:    blog.Title,
		Content:  blog.Content,
	}
}
