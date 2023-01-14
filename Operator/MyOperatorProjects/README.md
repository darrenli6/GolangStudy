
# init

```
kubebuilder  init  \ 
--domain ai-lijia.com  \ 
--repo=github.com/darrenli6/application-operator   \
--owner "Darren Li" 


go mod tidy 



```
 
# create api

```
kubebuilder create api  \
--group app  \
--version v1  \
--kind Application 
```




# make manifests

命令生成clusterRole和CRD配置

