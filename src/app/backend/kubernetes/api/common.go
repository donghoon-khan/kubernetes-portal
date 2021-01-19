package api

import (
	"crypto/rand"
	"fmt"
	"strings"

	v1 "k8s.io/api/authorization/v1"
)

func ToSelfSubjectAccessReview(namespace, name, resource, verb string) *v1.SelfSubjectAccessReview {
	return &v1.SelfSubjectAccessReview{
		Spec: v1.SelfSubjectAccessReviewSpec{
			ResourceAttributes: &v1.ResourceAttributes{
				Namespace: namespace,
				Name:      name,
				Resource:  fmt.Sprintf("%ss", strings.ToLower(resource)),
				Verb:      strings.ToLower(verb),
			},
		},
	}
}

func GenerateCSRFKey() string {
	bytes := make([]byte, 256)
	_, err := rand.Read(bytes)
	if err != nil {
		panic("could not generate csrf key")
	}

	return string(bytes)
}
