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
			},
			"service_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Service to link the kafka topic to",
			},
			"topic": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Topic name",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username",
			},
			"permission": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Permission to create",
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
