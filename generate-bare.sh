set -x
set -e

mkdir -p $1

fulldir=platforms/$1/$2
mkdir -p $fulldir

pushd $fulldir
mkdir -p problem test-cases solutions
touch solutions/solution.go solutions/solution.py
code  solutions/solution.go solutions/solution.py
popd

nautilus $fulldir

