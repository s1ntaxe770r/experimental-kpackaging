cluster:
	kind create cluster --name kpack-playground --image kindest/node:v1.27.2 --wait 30s 

# install kpack 
install:
	kubectl apply -f ./kpack.yaml

# setup kpack 
setup:
	kubectl apply -f ./service-account.yaml
	kubectl apply -f ./store.yaml 
	sleep 15
	kubectl apply -f ./stack.yaml
	sleep 15
	kubectl apply -f ./builder.yaml
