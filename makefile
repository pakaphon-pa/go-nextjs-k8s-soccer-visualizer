k8s-deploy:
	@echo "k8s deployment"
	kubectl apply -f deploy/namespace.yml
	kubectl apply -f deploy/
	
k8s-clear:
	@echo "clear k8s deployment"
	kubectl delete all --all -n learning
	kubectl delete pvc --all -n learning
	kubectl delete pv --all -n learning
	kubectl delete configmap --all -n learning
	kubectl delete secret --all -n learning
	kubectl delete ingress --all -n learning
	kubectl delete namespace learning

colima-reset:
	@echo "💣 Resetting Colima Kubernetes VM..."
	colima stop
	colima delete
	colima start --kubernetes --memory 4 --cpu 2 --disk 20