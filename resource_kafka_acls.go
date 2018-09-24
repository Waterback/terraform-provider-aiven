package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jelmersnoeck/aiven"
)

func resourceKafkaAcls() *schema.Resource {
	return &schema.Resource{
		Create: resourceKafkaAclCreate,

		Schema: map[string]*schema.Schema{
			"project": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Project to link the kafka topic to",
				ForceNew:    true,
			},
			"service_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Service to link the kafka topic to",
				ForceNew:    true,
			},
			"topic": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Topic name",
				ForceNew:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username",
				ForceNew:    true,
			},
			"permission": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Permission to create",
				ForceNew:    true,
			},
		},
	}
}

func resourceKafkaAclCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*aiven.Client)

	project := d.Get("project").(string)
	serviceName := d.Get("service_name").(string)
	topic := d.Get("topic").(string)
	permission := d.Get("replication").(string)
	username := d.Get("partitions").(string)

	acls, err := client.KafkaAcls.Create(
		project,
		serviceName,
		aiven.CreateKafkaAclRequest{
			Permission: &permission,
			Topic:      &topic,
			Username:   &username,
		},
	)
	if err != nil {
		d.SetId("")
		return err
	}

	d.SetId(project + "/" + serviceName + "/acl" + *acls[0].Id)

	return nil
}
