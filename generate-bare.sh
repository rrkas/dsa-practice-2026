set -x
set -e

mkdir -p $1

fulldir=$1/$2
mkdir -p $fulldir

pushd $fulldir
mkdir -p problem test-cases
touch solution.go solution.py
code  solution.go solution.py
popd

nautilus $fulldir

