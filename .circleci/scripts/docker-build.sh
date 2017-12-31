#!/bin/bash -e

REGISTRY=ernoaapa
PLATFORMS="linux/amd64,linux/arm64"
GIT_HASH=$(git describe --tags --always --dirty)


for bin in eliotd; do
  image="${REGISTRY}/${bin}"
  
  for osarch in ${PLATFORMS//,/ }; do
    os="${osarch%%/*}"
    arch="${osarch#*/}"
    echo "Building container for: $bin $os $arch"

    sed \
	    -e "s|ARG_BIN|${bin}|g" \
			-e "s|ARG_OS|${os}|g" \
			-e "s|ARG_ARCH|${arch}|g" \
	    Dockerfile.tmpl > .dockerfile-${arch}

    version="${GIT_HASH}-${arch}"
    tag="${image}:${version}"
    echo "Build docker image ${tag}:"
    echo "---------------Dockerfile start---------------------"
    cat .dockerfile-${arch}
    echo "----------------Dockerfile end----------------------"

    docker build -t ${tag} -f .dockerfile-${arch} .
    
    # Test spin up amd64 container because running in circleCI
    if [ $arch == "amd64" ]; then
      echo "Test running the container by printing help text"
      docker run -it ${tag} -h
    fi
  done
done
