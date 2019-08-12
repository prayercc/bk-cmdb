package proc_server_test

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"configcenter/src/common/metadata"
	params "configcenter/src/common/paraparse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("service template test", func() {
	var categoryId, serviceTemplateId, moduleId, serviceId, serviceId1, processTemplateId, processId int64
	resMap := make(map[string]interface{}, 0)

	Describe("service template test", func() {
		It("create service category", func() {
			input := map[string]interface{}{
				"bk_parent_id": 0,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"name": "test10",
			}
			rsp, err := serviceClient.CreateServiceCategory(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := metadata.ServiceCategory{}
			json.Unmarshal(j, &data)
			categoryId = data.ID
		})

		It("search service category", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
			}
			rsp, err := serviceClient.SearchServiceCategory(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			resMap["service_category"] = j
		})

		It("create service template", func() {
			input := map[string]interface{}{
				"service_category_id": categoryId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"name": "st",
			}
			rsp, err := serviceClient.CreateServiceTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := metadata.ServiceTemplate{}
			json.Unmarshal(j, &data)
			Expect(data.Name).To(Equal("st"))
			Expect(data.ServiceCategoryID).To(Equal(categoryId))
			serviceTemplateId = data.ID
		})

		It("search service template", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"service_category_id": categoryId,
			}
			rsp, err := serviceClient.SearchServiceTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(ContainSubstring("\"count\":1"))
			Expect(j).To(ContainSubstring("\"name\":\"st\""))
			Expect(j).To(ContainSubstring(fmt.Sprintf("\"service_category_id\":%d", categoryId)))
			resMap["service_template"] = j
		})

		It("create service template with empty name", func() {
			input := map[string]interface{}{
				"service_category_id": categoryId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"name": "",
			}
			rsp, err := serviceClient.CreateServiceTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199006))
		})

		It("create service template with invalid service category", func() {
			input := map[string]interface{}{
				"service_category_id": 12345,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"name": "st1",
			}
			rsp, err := serviceClient.CreateServiceTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199006))
		})

		It("search service template", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"service_category_id": categoryId,
			}
			rsp, err := serviceClient.SearchServiceTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(Equal(resMap["service_template"]))
		})

		It("delete service category with template", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"id": categoryId,
			}
			rsp, err := serviceClient.DeleteServiceCategory(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199056))
		})

		It("search service category", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
			}
			rsp, err := serviceClient.SearchServiceCategory(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(Equal(resMap["service_category"]))
		})

		It("create module with template", func() {
			input := map[string]interface{}{
				"bk_module_name":      "test12345",
				"bk_parent_id":        setId,
				"service_category_id": categoryId,
				"service_template_id": serviceTemplateId,
			}
			rsp, err := instClient.CreateModule(context.Background(), strconv.FormatInt(bizId, 10), strconv.FormatInt(setId, 10), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			moduleId = int64(rsp.Data["bk_module_id"].(float64))
		})

		It("search module", func() {
			input := &params.SearchParams{
				Condition: map[string]interface{}{},
				Page: map[string]interface{}{
					"sort": "id",
				},
			}
			rsp, err := instClient.SearchModule(context.Background(), "0", strconv.FormatInt(bizId, 10), strconv.FormatInt(setId, 10), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(ContainSubstring("\"bk_module_name\":\"st\""))
			Expect(j).To(ContainSubstring(fmt.Sprintf("\"service_template_id\":%d", serviceTemplateId)))
			Expect(j).To(ContainSubstring(fmt.Sprintf("\"service_category_id\":%d", categoryId)))
			resMap["module"] = j
		})

		It("create module with invalid template", func() {
			input := map[string]interface{}{
				"bk_module_name":      "12345",
				"bk_parent_id":        setId,
				"service_category_id": categoryId,
				"service_template_id": 10000,
			}
			rsp, err := instClient.CreateModule(context.Background(), strconv.FormatInt(bizId, 10), strconv.FormatInt(setId, 10), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199006))
		})

		It("create module with unmatch category and template", func() {
			input := map[string]interface{}{
				"bk_module_name":      "123",
				"bk_parent_id":        setId,
				"service_category_id": 2,
				"service_template_id": serviceTemplateId,
			}
			rsp, err := instClient.CreateModule(context.Background(), strconv.FormatInt(bizId, 10), strconv.FormatInt(setId, 10), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1108036))
		})

		It("create module with same template", func() {
			input := map[string]interface{}{
				"bk_module_name":      "test1234567",
				"bk_parent_id":        setId,
				"service_category_id": categoryId,
				"service_template_id": serviceTemplateId,
			}
			rsp, err := instClient.CreateModule(context.Background(), strconv.FormatInt(bizId, 10), strconv.FormatInt(setId, 10), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199014))
		})

		It("search module", func() {
			input := &params.SearchParams{
				Condition: map[string]interface{}{},
				Page: map[string]interface{}{
					"sort": "id",
				},
			}
			rsp, err := instClient.SearchModule(context.Background(), "0", strconv.FormatInt(bizId, 10), strconv.FormatInt(setId, 10), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(Equal(resMap["module"]))
		})

		It("update module with template", func() {
			input := map[string]interface{}{
				"bk_module_name":      "TEST",
				"service_category_id": 2,
				"service_template_id": 1000,
			}
			rsp, err := instClient.UpdateModule(context.Background(), strconv.FormatInt(bizId, 10), strconv.FormatInt(setId, 10), strconv.FormatInt(moduleId, 10), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
		})

		It("search module", func() {
			input := &params.SearchParams{
				Condition: map[string]interface{}{},
				Page: map[string]interface{}{
					"sort": "id",
				},
			}
			rsp, err := instClient.SearchModule(context.Background(), "0", strconv.FormatInt(bizId, 10), strconv.FormatInt(setId, 10), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(ContainSubstring("\"bk_module_name\":\"st\""))
			Expect(j).To(ContainSubstring(fmt.Sprintf("\"service_template_id\":%d", serviceTemplateId)))
			Expect(j).To(ContainSubstring(fmt.Sprintf("\"service_category_id\":%d", categoryId)))
		})

		It("delete service template with module", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"service_template_id": serviceTemplateId,
			}
			rsp, err := serviceClient.DeleteServiceTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199056))
		})

		It("search service template", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"service_category_id": categoryId,
			}
			rsp, err := serviceClient.SearchServiceTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(Equal(resMap["service_template"]))
		})
	})

	Describe("service instance test", func() {
		It("create service instance with template", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"bk_module_id": moduleId,
				"instances": []map[string]interface{}{
					{
						"bk_host_id": hostId1,
						"processes": []map[string]interface{}{
							{
								"process_info": map[string]interface{}{
									"bk_func_name":    "p1",
									"bk_process_name": "p1",
								},
							},
						},
					},
				},
			}
			rsp, err := serviceClient.CreateServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			serviceId = int64(rsp.Data.([]interface{})[0].(float64))
		})

		It("search service instance", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(ContainSubstring(fmt.Sprintf("\"id\":%d", serviceId)))
			resMap["service_instance"] = j
		})

		It("clone service instance to source host", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"bk_module_id": moduleId,
				"instances": []map[string]interface{}{
					{
						"bk_host_id": hostId1,
						"processes": []map[string]interface{}{
							{
								"process_info": map[string]interface{}{
									"bk_func_name":    "p1",
									"bk_process_name": "p1",
								},
							},
						},
					},
				},
			}
			rsp, err := serviceClient.CreateServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1113016))
		})

		It("search service instance", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(Equal(resMap["service_instance"]))
		})
	})

	Describe("process template test", func() {
		It("create process template", func() {
			input := map[string]interface{}{
				"service_template_id": serviceTemplateId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"processes": []map[string]interface{}{
					{
						"spec": map[string]interface{}{
							"bk_func_name": map[string]interface{}{
								"value":            "p1",
								"as_default_value": true,
							},
							"bk_process_name": map[string]interface{}{
								"value":            "p1",
								"as_default_value": true,
							},
							"bk_start_param_regex": map[string]interface{}{
								"value":            "123",
								"as_default_value": false,
							},
						},
					},
				},
			}
			rsp, err := processClient.CreateProcessTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			processTemplateId = int64(rsp.Data.([]interface{})[0].(float64))
		})

		It("search process template", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"service_template_id": serviceTemplateId,
			}
			rsp, err := processClient.SearchProcessTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(ContainSubstring("\"bk_func_name\":{\"as_default_value\":true,\"value\":\"p1\"}"))
			Expect(j).To(ContainSubstring("\"bk_process_name\":{\"as_default_value\":true,\"value\":\"p1\"}"))
			Expect(j).To(ContainSubstring("\"bk_start_param_regex\":{\"as_default_value\":false,\"value\":\"123\"}"))
			Expect(j).To(ContainSubstring(fmt.Sprintf("\"id\":%d", processTemplateId)))
			resMap["process_template"] = j
		})

		It("create process template with same name", func() {
			input := map[string]interface{}{
				"service_template_id": serviceTemplateId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"processes": []map[string]interface{}{
					{
						"spec": map[string]interface{}{
							"bk_func_name": map[string]interface{}{
								"value":            "p123",
								"as_default_value": true,
							},
							"bk_process_name": map[string]interface{}{
								"value":            "p1",
								"as_default_value": true,
							},
						},
					},
				},
			}
			rsp, err := processClient.CreateProcessTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1113019))
		})

		It("create process template with same bk_func_name and bk_start_param_regex", func() {
			input := map[string]interface{}{
				"service_template_id": serviceTemplateId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"processes": []map[string]interface{}{
					{
						"spec": map[string]interface{}{
							"bk_func_name": map[string]interface{}{
								"value":            "p1",
								"as_default_value": true,
							},
							"bk_process_name": map[string]interface{}{
								"value":            "p123",
								"as_default_value": true,
							},
							"bk_start_param_regex": map[string]interface{}{
								"value":            "123",
								"as_default_value": false,
							},
						},
					},
				},
			}
			rsp, err := processClient.CreateProcessTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1113020))
		})

		It("create process template with empty name", func() {
			input := map[string]interface{}{
				"service_template_id": serviceTemplateId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"processes": []map[string]interface{}{
					{
						"spec": map[string]interface{}{
							"bk_func_name": map[string]interface{}{
								"value":            "",
								"as_default_value": true,
							},
							"bk_process_name": map[string]interface{}{
								"value":            "",
								"as_default_value": true,
							},
						},
					},
				},
			}
			rsp, err := processClient.CreateProcessTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199006))
		})

		It("create process template with invalid service template", func() {
			input := map[string]interface{}{
				"service_template_id": 10000,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"processes": []map[string]interface{}{
					{
						"spec": map[string]interface{}{
							"bk_func_name": map[string]interface{}{
								"value":            "123",
								"as_default_value": true,
							},
							"bk_process_name": map[string]interface{}{
								"value":            "123",
								"as_default_value": true,
							},
						},
					},
				},
			}
			rsp, err := processClient.CreateProcessTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199006))
		})

		It("search process template", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"service_template_id": serviceTemplateId,
			}
			rsp, err := processClient.SearchProcessTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(Equal(resMap["process_template"]))
		})

		It("clone service instance to other host", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"bk_module_id": moduleId,
				"instances": []map[string]interface{}{
					{
						"bk_host_id": hostId2,
						"processes": []map[string]interface{}{
							{
								"process_info": map[string]interface{}{
									"bk_func_name":    "p2",
									"bk_process_name": "p3",
								},
							},
						},
					},
				},
			}
			rsp, err := serviceClient.CreateServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			serviceId1 = int64(rsp.Data.([]interface{})[0].(float64))
		})

		It("search service instance", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(ContainSubstring(fmt.Sprintf("\"id\":%d", serviceId)))
			Expect(j).To(ContainSubstring(fmt.Sprintf("\"id\":%d", serviceId1)))
			resMap["service_instance"] = j
		})

		It("search process instance", func() {
			input := map[string]interface{}{
				"service_instance_id": serviceId1,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
			}
			rsp, err := processClient.SearchProcessInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := []metadata.ProcessInstance{}
			json.Unmarshal(j, &data)
			Expect(len(data)).To(Equal(1))
			Expect(data[0].Property["bk_process_name"]).To(Equal("p1"))
			Expect(data[0].Property["bk_func_name"]).To(Equal("p1"))
			Expect(data[0].Property["bk_start_param_regex"]).To(Equal("123"))
			Expect(data[0].Relation.HostID).To(Equal(hostId2))
			processId = int64(data[0].Property["bk_process_id"].(float64))
		})

		It("create process instance", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"service_instance_id": serviceId,
				"processes": []map[string]interface{}{
					{
						"process_info": map[string]interface{}{
							"bk_process_name": "p2",
							"bk_func_name":    "p2",
						},
					},
				},
			}
			rsp, err := processClient.CreateProcessInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1108035))
		})

		It("udpate process instance", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"processes": []map[string]interface{}{
					{
						"bk_func_name":         "p3",
						"bk_process_name":      "p3",
						"bk_start_param_regex": "1234",
						"bk_process_id":        processId,
					},
				},
			}
			rsp, err := processClient.UpdateProcessInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
		})

		It("search process instance", func() {
			input := map[string]interface{}{
				"service_instance_id": serviceId1,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
			}
			rsp, err := processClient.SearchProcessInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := []metadata.ProcessInstance{}
			json.Unmarshal(j, &data)
			Expect(len(data)).To(Equal(1))
			Expect(data[0].Property["bk_process_name"]).To(Equal("p1"))
			Expect(data[0].Property["bk_func_name"]).To(Equal("p1"))
			Expect(data[0].Property["bk_start_param_regex"]).To(Equal("1234"))
			Expect(data[0].Relation.HostID).To(Equal(hostId2))
		})

		It("update process template", func() {
			input := map[string]interface{}{
				"process_template_id": processTemplateId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"process_property": map[string]interface{}{
					"bk_func_name": map[string]interface{}{
						"value":            "123",
						"as_default_value": false,
					},
					"bk_process_name": map[string]interface{}{
						"value":            "123",
						"as_default_value": false,
					},
					"bk_start_param_regex": map[string]interface{}{
						"value":            "123456",
						"as_default_value": true,
					},
				},
			}
			rsp, err := processClient.UpdateProcessTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(ContainSubstring("\"bk_func_name\":{\"as_default_value\":true,\"value\":\"p1\"}"))
			Expect(j).To(ContainSubstring("\"bk_process_name\":{\"as_default_value\":true,\"value\":\"p1\"}"))
			Expect(j).To(ContainSubstring("\"bk_start_param_regex\":{\"as_default_value\":true,\"value\":\"123456\"}"))
		})

		It("search process template", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"service_template_id": serviceTemplateId,
			}
			rsp, err := processClient.SearchProcessTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(ContainSubstring("\"bk_func_name\":{\"as_default_value\":true,\"value\":\"p1\"}"))
			Expect(j).To(ContainSubstring("\"bk_process_name\":{\"as_default_value\":true,\"value\":\"p1\"}"))
			Expect(j).To(ContainSubstring("\"bk_start_param_regex\":{\"as_default_value\":true,\"value\":\"123456\"}"))
			resMap["process_template"] = j
		})
	})

	Describe("update template test", func() {
		It("update service template", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"id":                  serviceTemplateId,
				"service_category_id": 2,
				"name":                "abcdefg",
			}
			rsp, err := serviceClient.UpdateServiceTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := metadata.ServiceTemplate{}
			json.Unmarshal(j, &data)
			Expect(data.Name).To(Equal("st"))
			Expect(data.ServiceCategoryID).To(Equal(int64(2)))
		})

		It("search service template", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"service_category_id": 2,
			}
			rsp, err := serviceClient.SearchServiceTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(ContainSubstring("\"count\":1"))
			Expect(j).To(ContainSubstring("\"name\":\"st\""))
			Expect(j).To(ContainSubstring("\"service_category_id\":2"))
			resMap["service_template"] = j
		})

		It("update service template with invalid service category", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"id":                  serviceTemplateId,
				"service_category_id": 100000,
				"name":                "abcdefg",
			}
			rsp, err := serviceClient.UpdateServiceTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199006))
		})

		It("search service template", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"service_category_id": 2,
			}
			rsp, err := serviceClient.SearchServiceTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(Equal(resMap["service_template"]))
		})

		It("compare service instance and template after add and change process template", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"bk_module_id":        moduleId,
				"service_template_id": serviceTemplateId,
			}
			rsp, err := serviceClient.DiffServiceInstanceWithTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := new(metadata.ProcessTemplateWithInstancesDifference)
			json.Unmarshal(j, data)
			Expect(len(data.Removed)).To(Equal(0))
			Expect(len(data.Unchanged)).To(Equal(0))
			Expect(len(data.Added)).To(Equal(1))
			Expect(len(data.Changed)).To(Equal(1))
			Expect(data.Added[0].ProcessTemplateID).To(Equal(processTemplateId))
			Expect(data.Added[0].ServiceInstanceCount).To(Equal(1))
			Expect(data.Added[0].ServiceInstances[0].ServiceInstance.ID).To(Equal(serviceId))
			Expect(len(data.Changed)).To(Equal(1))
			Expect(data.Changed[0].ProcessTemplateID).To(Equal(processTemplateId))
			Expect(data.Changed[0].ServiceInstanceCount).To(Equal(1))
			Expect(data.Changed[0].ServiceInstances[0].ServiceInstance.ID).To(Equal(serviceId1))
			Expect(len(data.Changed[0].ServiceInstances[0].ChangedAttributes)).To(Equal(1))
			Expect(data.Changed[0].ServiceInstances[0].ChangedAttributes[0].PropertyID).To(Equal("bk_start_param_regex"))
			Expect(data.Changed[0].ServiceInstances[0].ChangedAttributes[0].PropertyValue).To(Equal("1234"))
			j, err = json.Marshal(data.Changed[0].ServiceInstances[0].ChangedAttributes[0].TemplatePropertyValue)
			Expect(j).To(Equal("{\"as_default_value\":true,\"value\":\"123456\"}"))
		})

		It("sync service instance and template after add and change process template", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"bk_module_id":        moduleId,
				"service_template_id": serviceTemplateId,
			}
			rsp, err := serviceClient.SyncServiceInstanceByTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
		})

		It("search module", func() {
			input := &params.SearchParams{
				Condition: map[string]interface{}{},
				Page: map[string]interface{}{
					"sort": "id",
				},
			}
			rsp, err := instClient.SearchModule(context.Background(), "0", strconv.FormatInt(bizId, 10), strconv.FormatInt(setId, 10), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(ContainSubstring("\"bk_module_name\":\"st\""))
			Expect(j).To(ContainSubstring(fmt.Sprintf("\"service_template_id\":%d", serviceTemplateId)))
			Expect(j).To(ContainSubstring("\"service_category_id\":2"))
		})

		It("search process instance", func() {
			input := map[string]interface{}{
				"service_instance_id": serviceId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
			}
			rsp, err := processClient.SearchProcessInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := []metadata.ProcessInstance{}
			json.Unmarshal(j, &data)
			Expect(len(data)).To(Equal(1))
			Expect(data[0].Property["bk_process_name"]).To(Equal("p1"))
			Expect(data[0].Property["bk_func_name"]).To(Equal("p1"))
			Expect(data[0].Property["bk_start_param_regex"]).To(Equal("123456"))
			Expect(data[0].Relation.HostID).To(Equal(hostId1))
		})

		It("udpate process instance", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"processes": []map[string]interface{}{
					{
						"bk_func_name":         "p3",
						"bk_process_name":      "p3",
						"bk_start_param_regex": "1234",
						"bk_process_id":        processId,
					},
				},
			}
			rsp, err := processClient.UpdateProcessInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
		})

		It("delete process instance", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"process_instance_ids": []int64{
					processId,
				},
			}
			rsp, err := processClient.DeleteProcessInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1113022))
		})

		It("search process instance", func() {
			input := map[string]interface{}{
				"service_instance_id": serviceId1,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
			}
			rsp, err := processClient.SearchProcessInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := []metadata.ProcessInstance{}
			json.Unmarshal(j, &data)
			Expect(len(data)).To(Equal(1))
			Expect(data[0].Property["bk_process_name"]).To(Equal("p1"))
			Expect(data[0].Property["bk_func_name"]).To(Equal("p1"))
			Expect(data[0].Property["bk_start_param_regex"]).To(Equal("123456"))
		})
	})

	Describe("service instance label test", func() {
		It("service instance add labels", func() {
			input := map[string]interface{}{
				"labels": map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
				},
				"instance_ids": []int64{
					serviceId,
					serviceId1,
				},
			}
			rsp, err := serviceClient.ServiceInstanceAddLabels(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
		})

		It("service instance add and edit labels", func() {
			input := map[string]interface{}{
				"labels": map[string]interface{}{
					"key2": "value",
					"key3": "value3",
				},
				"instance_ids": []int64{
					serviceId1,
				},
			}
			rsp, err := serviceClient.ServiceInstanceAddLabels(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
		})

		It("search service instance", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			resMap["service_instance"] = j
			data := new(metadata.MultipleServiceInstance)
			json.Unmarshal(j, &data)
			Expect(data.Count).To(Equal(uint64(2)))
			Expect(len(data.Info[0].Labels)).To(Equal(2))
			Expect(data.Info[0].Labels["key1"]).To(Equal("value1"))
			Expect(data.Info[0].Labels["key2"]).To(Equal("value2"))
			Expect(len(data.Info[1].Labels)).To(Equal(3))
			Expect(data.Info[1].Labels["key1"]).To(Equal("value1"))
			Expect(data.Info[1].Labels["key2"]).To(Equal("value"))
			Expect(data.Info[1].Labels["key3"]).To(Equal("value3"))
		})

		It("service instance add labels with empty key values", func() {
			input := map[string]interface{}{
				"labels": map[string]interface{}{
					"":     "value1",
					"key1": "",
					"key4": "value4",
				},
				"instance_ids": []int64{
					serviceId,
					serviceId1,
				},
			}
			rsp, err := serviceClient.ServiceInstanceAddLabels(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199006))
		})

		It("service instance add labels with invalid instance id", func() {
			input := map[string]interface{}{
				"labels": map[string]interface{}{
					"key5": "value5",
				},
				"instance_ids": []int64{
					serviceId,
					10000,
					serviceId1,
				},
			}
			rsp, err := serviceClient.ServiceInstanceAddLabels(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199017))
		})

		It("search module service instances labels", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
			}
			rsp, err := serviceClient.ServiceInstanceFindLabels(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := make(map[string][]string)
			json.Unmarshal(j, &data)
			Expect(len(data)).To(Equal(3))
			Expect(data["key1"]).To(ConsistOf("value1"))
			Expect(data["key2"]).To(ConsistOf("value2", "value"))
			Expect(data["key3"]).To(ConsistOf("value3"))
		})

		It("search service instance", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			Expect(resMap["service_instance"]).To(Equal(j))
		})

		It("search service instance without key", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"selectors": []map[string]interface{}{
					{
						"key":      "key3",
						"operator": "!",
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := new(metadata.MultipleServiceInstance)
			json.Unmarshal(j, &data)
			Expect(data.Count).To(Equal(uint64(1)))
			Expect(data.Info[0].ID).To(Equal(serviceId))
		})

		It("search service instance without key with values", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"selectors": []map[string]interface{}{
					{
						"key":      "key3",
						"operator": "!",
						"values": []string{
							"123",
						},
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199006))
		})

		It("search service instance exists key", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"selectors": []map[string]interface{}{
					{
						"key":      "key3",
						"operator": "exists",
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := new(metadata.MultipleServiceInstance)
			json.Unmarshal(j, &data)
			Expect(data.Count).To(Equal(uint64(1)))
			Expect(data.Info[0].ID).To(Equal(serviceId1))
		})

		It("search service instance exists key with values", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"selectors": []map[string]interface{}{
					{
						"key":      "key3",
						"operator": "exists",
						"values": []string{
							"123",
						},
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199006))
		})

		It("search service instance with equal key value", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"selectors": []map[string]interface{}{
					{
						"key":      "key1",
						"operator": "=",
						"values": []string{
							"value1",
						},
					},
					{
						"key":      "key2",
						"operator": "=",
						"values": []string{
							"value2",
						},
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := new(metadata.MultipleServiceInstance)
			json.Unmarshal(j, &data)
			Expect(data.Count).To(Equal(uint64(1)))
			Expect(data.Info[0].ID).To(Equal(serviceId))
		})

		It("search service instance with equal key zero value", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"selectors": []map[string]interface{}{
					{
						"key":      "key1",
						"operator": "=",
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199006))
		})

		It("search service instance with equal key many values", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"selectors": []map[string]interface{}{
					{
						"key":      "key1",
						"operator": "=",
						"values": []string{
							"value1",
							"value2",
						},
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199006))
		})

		It("search service instance with not equal key value", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"selectors": []map[string]interface{}{
					{
						"key":      "key1",
						"operator": "!=",
						"values": []string{
							"value2",
						},
					},
					{
						"key":      "key2",
						"operator": "!=",
						"values": []string{
							"value",
						},
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := new(metadata.MultipleServiceInstance)
			json.Unmarshal(j, &data)
			Expect(data.Count).To(Equal(uint64(1)))
			Expect(data.Info[0].ID).To(Equal(serviceId))
		})

		It("search service instance with not equal key zero value", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"selectors": []map[string]interface{}{
					{
						"key":      "key1",
						"operator": "!=",
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199006))
		})

		It("search service instance with not equal key many values", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"selectors": []map[string]interface{}{
					{
						"key":      "key1",
						"operator": "!=",
						"values": []string{
							"value1",
							"value2",
						},
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199006))
		})

		It("search service instance with value in values", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"selectors": []map[string]interface{}{
					{
						"key":      "key1",
						"operator": "in",
						"values": []string{
							"value1",
						},
					},
					{
						"key":      "key2",
						"operator": "in",
						"values": []string{
							"value",
							"value2",
						},
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := new(metadata.MultipleServiceInstance)
			json.Unmarshal(j, &data)
			Expect(data.Count).To(Equal(uint64(2)))
			Expect(data.Info[0].ID).To(Equal(serviceId))
			Expect(data.Info[1].ID).To(Equal(serviceId1))
		})

		It("search service instance with value in zero values", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"selectors": []map[string]interface{}{
					{
						"key":      "key1",
						"operator": "in",
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199006))
		})

		It("search service instance with value not in values", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"selectors": []map[string]interface{}{
					{
						"key":      "key3",
						"operator": "notin",
						"values": []string{
							"value",
						},
					},
					{
						"key":      "key1",
						"operator": "notin",
						"values": []string{
							"value",
							"value2",
						},
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := new(metadata.MultipleServiceInstance)
			json.Unmarshal(j, &data)
			Expect(data.Count).To(Equal(uint64(2)))
			Expect(data.Info[0].ID).To(Equal(serviceId))
			Expect(data.Info[1].ID).To(Equal(serviceId1))
		})

		It("search service instance with value not in zero values", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"selectors": []map[string]interface{}{
					{
						"key":      "key1",
						"operator": "notin",
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199006))
		})

		It("search service instance with invalid operator", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"selectors": []map[string]interface{}{
					{
						"key":      "key1",
						"operator": "123",
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199006))
		})

		It("search service instance with empty key", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"selectors": []map[string]interface{}{
					{
						"key":      "",
						"operator": "exists",
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199006))
		})

		It("search service instance with no matching data", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"selectors": []map[string]interface{}{
					{
						"key":      "key1",
						"operator": "!",
					},
					{
						"key":      "key3",
						"operator": "exists",
					},
					{
						"key":      "key3",
						"operator": "notin",
						"values": []string{
							"value",
						},
					},
					{
						"key":      "key2",
						"operator": "!=",
						"values": []string{
							"value3",
						},
					},
					{
						"key":      "key2",
						"operator": "=",
						"values": []string{
							"value2",
						},
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			Expect(j).To(ContainSubstring("\"count\":0"))
		})

		It("service instance remove labels", func() {
			input := map[string]interface{}{
				"keys": []string{
					"key1",
					"",
					"key3",
				},
				"instance_ids": []int64{
					serviceId,
					serviceId1,
				},
			}
			rsp, err := serviceClient.ServiceInstanceRemoveLabels(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
		})

		It("service instance remove labels with invalid service instance id", func() {
			input := map[string]interface{}{
				"keys": []string{
					"key2",
				},
				"instance_ids": []int64{
					serviceId,
					100000,
					serviceId1,
				},
			}
			rsp, err := serviceClient.ServiceInstanceRemoveLabels(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(1199017))
		})

		It("search service instance", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := new(metadata.MultipleServiceInstance)
			json.Unmarshal(j, &data)
			Expect(data.Count).To(Equal(uint64(2)))
			Expect(len(data.Info[0].Labels)).To(Equal(1))
			Expect(data.Info[0].Labels["key2"]).To(Equal("value2"))
			Expect(len(data.Info[1].Labels)).To(Equal(1))
			Expect(data.Info[1].Labels["key2"]).To(Equal("value"))
		})
	})

	Describe("removal test", func() {
		It("remove process template", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"process_templates": []int64{
					processTemplateId,
				},
			}
			rsp, err := processClient.DeleteProcessTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
		})

		It("search process template", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"service_template_id": serviceTemplateId,
			}
			rsp, err := processClient.SearchProcessTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp)
			Expect(j).To(ContainSubstring("\"count\":0"))
		})

		It("compare service instance and template after add and change process template", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"bk_module_id":        moduleId,
				"service_template_id": serviceTemplateId,
			}
			rsp, err := serviceClient.DiffServiceInstanceWithTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := new(metadata.ProcessTemplateWithInstancesDifference)
			json.Unmarshal(j, data)
			Expect(len(data.Removed)).To(Equal(1))
			Expect(len(data.Unchanged)).To(Equal(0))
			Expect(len(data.Added)).To(Equal(0))
			Expect(len(data.Changed)).To(Equal(0))
			Expect(data.Removed[0].ProcessTemplateID).To(Equal(processTemplateId))
			Expect(data.Removed[0].ServiceInstanceCount).To(Equal(2))
			Expect(data.Removed[0].ServiceInstances[0].ServiceInstance.ID).To(Equal(serviceId))
			Expect(data.Removed[0].ServiceInstances[1].ServiceInstance.ID).To(Equal(serviceId1))
		})

		It("remove service instance with template with process", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"service_instance_ids": []int64{
					serviceId1,
				},
			}
			rsp, err := serviceClient.DeleteServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
		})

		It("search service instance", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := new(metadata.MultipleServiceInstance)
			json.Unmarshal(j, &data)
			Expect(data.Count).To(Equal(uint64(1)))
			Expect(data.Info[0].ID).To(Equal(serviceId))
		})

		It("sync service instance and template after add and change process template", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"bk_module_id":        moduleId,
				"service_template_id": serviceTemplateId,
			}
			rsp, err := serviceClient.SyncServiceInstanceByTemplate(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
		})

		It("search process instance", func() {
			input := map[string]interface{}{
				"service_instance_id": serviceId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
			}
			rsp, err := processClient.SearchProcessInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := []metadata.ProcessInstance{}
			json.Unmarshal(j, &data)
			Expect(len(data)).To(Equal(0))
		})

		It("remove service instance with template without process", func() {
			input := map[string]interface{}{
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
				"service_instance_ids": []int64{
					serviceId,
				},
			}
			rsp, err := serviceClient.DeleteServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
		})

		It("search service instance", func() {
			input := map[string]interface{}{
				"bk_module_id": moduleId,
				"metadata": map[string]interface{}{
					"label": map[string]interface{}{
						"bk_biz_id": strconv.FormatInt(bizId, 10),
					},
				},
			}
			rsp, err := serviceClient.SearchServiceInstance(context.Background(), header, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			j, err := json.Marshal(rsp.Data)
			data := new(metadata.MultipleServiceInstance)
			json.Unmarshal(j, &data)
			Expect(data.Count).To(Equal(uint64(0)))
		})

		// unbind service template on module is prohibited
		// It("unbind service template on module", func() {
		// 	input := map[string]interface{}{
		// 		"metadata": map[string]interface{}{
		// 			"label": map[string]interface{}{
		// 				"bk_biz_id": strconv.FormatInt(bizId, 10),
		// 			},
		// 		},
		// 		"bk_module_id": moduleId,
		// 	}
		// 	rsp, err := serviceClient.RemoveTemplateBindingOnModule(context.Background(), header, input)
		// 	Expect(err).NotTo(HaveOccurred())
		// 	Expect(rsp.Result).To(Equal(true))
		// })

		// It("search module", func() {
		// 	input := &params.SearchParams{
		// 		Condition: map[string]interface{}{},
		// 		Page: map[string]interface{}{
		// 			"sort": "id",
		// 		},
		// 	}
		// 	rsp, err := instClient.SearchModule(context.Background(), "0", strconv.FormatInt(bizId, 10), strconv.FormatInt(setId, 10), header, input)
		// 	Expect(err).NotTo(HaveOccurred())
		// 	Expect(rsp.Result).To(Equal(true))
		// 	j, err := json.Marshal(rsp)
		// 	Expect(j).To(ContainSubstring("\"service_template_id\":0"))
		// })

		// It("delete service template", func() {
		// 	input := map[string]interface{}{
		// 		"metadata": map[string]interface{}{
		// 			"label": map[string]interface{}{
		// 				"bk_biz_id": strconv.FormatInt(bizId, 10),
		// 			},
		// 		},
		// 		"service_template_id": serviceTemplateId,
		// 	}
		// 	rsp, err := serviceClient.DeleteServiceTemplate(context.Background(), header, input)
		// 	Expect(err).NotTo(HaveOccurred())
		// 	Expect(rsp.Result).To(Equal(true))
		// })

		// It("search service template", func() {
		// 	input := map[string]interface{}{
		// 		"metadata": map[string]interface{}{
		// 			"label": map[string]interface{}{
		// 				"bk_biz_id": strconv.FormatInt(bizId, 10),
		// 			},
		// 		},
		// 		"service_category_id": 2,
		// 	}
		// 	rsp, err := serviceClient.SearchServiceTemplate(context.Background(), header, input)
		// 	Expect(err).NotTo(HaveOccurred())
		// 	Expect(rsp.Result).To(Equal(true))
		// 	j, err := json.Marshal(rsp)
		// 	Expect(j).To(ContainSubstring("\"count\":0"))
		// })

		// It("delete service category", func() {
		// 	input := map[string]interface{}{
		// 		"metadata": map[string]interface{}{
		// 			"label": map[string]interface{}{
		// 				"bk_biz_id": strconv.FormatInt(bizId, 10),
		// 			},
		// 		},
		// 		"id": categoryId,
		// 	}
		// 	rsp, err := serviceClient.DeleteServiceCategory(context.Background(), header, input)
		// 	Expect(err).NotTo(HaveOccurred())
		// 	Expect(rsp.Result).To(Equal(true))
		// })

		// It("search service category", func() {
		// 	input := map[string]interface{}{
		// 		"metadata": map[string]interface{}{
		// 			"label": map[string]interface{}{
		// 				"bk_biz_id": strconv.FormatInt(bizId, 10),
		// 			},
		// 		},
		// 	}
		// 	rsp, err := serviceClient.SearchServiceCategory(context.Background(), header, input)
		// 	Expect(err).NotTo(HaveOccurred())
		// 	Expect(rsp.Result).To(Equal(true))
		// 	j, err := json.Marshal(rsp)
		// 	Expect(j).NotTo(ContainSubstring(fmt.Sprintf("\"id\":%d", categoryId)))
		// })

		// It("delete service category twice", func() {
		// 	input := map[string]interface{}{
		// 		"metadata": map[string]interface{}{
		// 			"label": map[string]interface{}{
		// 				"bk_biz_id": strconv.FormatInt(bizId, 10),
		// 			},
		// 		},
		// 		"id": categoryId,
		// 	}
		// 	rsp, err := serviceClient.DeleteServiceCategory(context.Background(), header, input)
		// 	Expect(err).NotTo(HaveOccurred())
		// 	Expect(rsp.Result).To(Equal(false))
		// 	Expect(rsp.Code).To(Equal(1199019))
		// })
	})
})
