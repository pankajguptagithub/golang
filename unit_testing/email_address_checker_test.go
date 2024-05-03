package linkedinlearning 

import "testing"

func TestIsLinkedInEmployee(t *testing.T){
	if IsLinkedInEmployee("a@lkedin.com") == false{
		t.Errorf("Expected false but got true")
	}

}
