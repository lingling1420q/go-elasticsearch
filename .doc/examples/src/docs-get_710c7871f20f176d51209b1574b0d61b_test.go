// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/get.asciidoc#L332>
//
// --------------------------------------------------------------------------------
// GET twitter/_doc/1?stored_fields=tags,counter
// --------------------------------------------------------------------------------

func Test_docs_get_710c7871f20f176d51209b1574b0d61b(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:710c7871f20f176d51209b1574b0d61b[]
	res, err := es.Get(
		"twitter",
		"1",
		es.Get.WithStoredFields("tags,counter"),
		es.Get.WithPretty(),
	)
	// end:710c7871f20f176d51209b1574b0d61b[]
	if err != nil {
		fmt.Println("Error getting the response:", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	fmt.Println(res)
}