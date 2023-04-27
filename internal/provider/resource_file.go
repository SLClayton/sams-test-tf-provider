package provider

import (
	"io/ioutil"
	"os"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceFileCreate,
		Read:   resourceFileRead,
		Delete: resourceFileDelete,
		Schema: map[string]*schema.Schema{
			"path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"content": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceFileCreate(d *schema.ResourceData, m interface{}) error {
	path := d.Get("path").(string)
	content := d.Get("content").(string)

	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return err
	}

	d.SetId(path)
	return resourceFileRead(d, m)
}

func resourceFileRead(d *schema.ResourceData, m interface{}) error {
	path := d.Id()
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	d.Set("path", path)
	d.Set("content", string(content))
	return nil
}

func resourceFileDelete(d *schema.ResourceData, m interface{}) error {
	path := d.Id()
	err := os.Remove(path) // Change this line
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
