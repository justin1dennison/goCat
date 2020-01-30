targets=${@-"darwin/amd64 linux/amd64  windows/amd64"}

for target in $targets; do
  os="$(echo $target | cut -d '/' -f1)"
  arch="$(echo $target | cut -d '/' -f2)"
  output="${os}_${arch}_goCat"

  echo "----> Building project for: $target"
  GOOS=$os GOARCH=$arch CGO_ENABLED=0 go build -o $output
done


echo "-----> Build is complete. List of files:"
cd $release_path
ls -al
