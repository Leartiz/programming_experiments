kubectl create ns new-deploy
kubectl create deployment nginx-deploy --image=nginx -n new-deploy

kubectl get deployments, rs, pods -n new-deploy
kubectl get deploy,rs,po -n new-deploy