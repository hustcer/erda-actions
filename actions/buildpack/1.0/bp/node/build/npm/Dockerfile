FROM {{BP_DOCKER_BASE_REGISTRY}}/erda-x/node:14

ARG PACKAGE_LOCK_DIR

COPY ${PACKAGE_LOCK_DIR} /app

ARG NODE_OPTIONS

ARG DEP_CMD="npm ci"

RUN cd /app && eval NODE_OPTIONS=${NODE_OPTIONS} NODE_ENV= ${DEP_CMD}

COPY /code /app

WORKDIR /app

ARG DICE_WORKSPACE

ARG WEBPACK_DLL_CONFIG="webpack_dll.config.js"

RUN \
    if [ -f "${WEBPACK_DLL_CONFIG}" ]; then \
        DICE_WORKSPACE=${DICE_WORKSPACE} npm run dll; \
    fi

RUN NODE_OPTIONS=${NODE_OPTIONS} DICE_WORKSPACE=${DICE_WORKSPACE} npm run build
