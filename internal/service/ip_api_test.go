package service

import (
	"context"
	"testing"
)

func TestIpApiService_GetLocation(t *testing.T) {
	s := NewIpApiService()
	resp, err := s.GetLocation(context.Background(), "1.1.1.1")
	if err != nil {
		t.Fatalf("GetLocation error: %v", err)
	}
	if resp.Status != "success" {
		t.Fatalf("GetLocation status error: %v", resp.Status)
	}
	t.Log(resp)
}
