
as_file="./out/a.s"
out_file="./out/a.out"
function test { 
  rm -f $as_file
  mkdir -p out
  echo "$1" | go run main.go > $as_file
  gcc -o $out_file $as_file
  
  ./out/a.out
  if [[ $? -eq 0 ]];then
    echo "ok"
  else
    echo "not ok"
    exit 1
  fi
}

test 0