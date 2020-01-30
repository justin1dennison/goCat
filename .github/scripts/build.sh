targets=${@-"darwin/amd64 linux/amd64  windows/amd64"}

for target in $targets; do
  os="$(echo $target | cut -d '/' -f1)"
  arch="$(echo $target | cut -d '/' -f2)"
  directory="${os}_${arch}"
  mkdir -p $directory
  output="${directory}/goCat"

  echo "----> Building project for: $target"
  GOOS=$os GOARCH=$arch CGO_ENABLED=0 go build -o $output
  zip -j "${directory}.zip" $directory
done


echo "-----> Build is complete. List of files:"
ls -al
