# YAML to JSON

You can`t install YQ (python package) on every environment, and this is simpler to audit, so for straightforward tasks, it is the best solution.


Example
```yaml
jobs:
   foo_job:     
     steps:
       - name: bar          
         uses: docker://alpine:3.8
```
```sh
$ yamlstring | yaml2json
{"jobs":{"foo_job":{"steps":[{"name":"bar","uses":"docker://alpine:3.8"}]}}}
# or with procesing 
$ yamlstring | yaml2json | jq -r '.jobs.foo_job.steps |  .[].uses'
docker://alpine:3.8
```
