#go mod tidy
#rm -f go.sum
go get -u github.com/Patrick-ring-motive/utils@0.0.3
go get -u github.com/Patrick-ring-motive/String@0.0.4
go get -u github.com/Patrick-ring-motive/httpne@0.0.7
go get -u github.com/Patrick-ring-motive/ione@0.0.23
go mod download github.com/Patrick-ring-motive/httpne

#go build -ldflags "-g" -gcflags="-B -v -std"  -o main main.go

cd api
go build -o handler vercel.go
chmod +x ./handler
cd main
go build -o indexgo index.go
chmod +x ./indexgo
cd ../..