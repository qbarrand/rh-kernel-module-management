FROM quay.io/quba/kmm-ci as source

FROM quay.io/edge-infrastructure/kernel-module-management-signimage:release-1.0 AS signimage

ARG SIGNROOT=/signroot

ARG UNSIGNED_KMOD_0
ARG UNSIGNED_KMOD_1
# ARG UNSIGNED_KMOD_2
# ARG ...

USER 0

RUN mkdir $SIGNROOT

COPY --from=source $UNSIGNED_KMOD_0 $SIGNROOT/$UNSIGNED_KMOD_0
COPY --from=source $UNSIGNED_KMOD_1 $SIGNROOT/$UNSIGNED_KMOD_1

RUN /usr/local/bin/sign-file sha256 /run/secrets/key.pem /run/secrets/cert.pem $SIGNROOT/$UNSIGNED_KMOD_0
RUN /usr/local/bin/sign-file sha256 /run/secrets/key.pem /run/secrets/cert.pem $SIGNROOT/$UNSIGNED_KMOD_1

FROM source

ARG SIGNROOT=/signroot
ARG UNSIGNED_KMOD_0
ARG UNSIGNED_KMOD_1
# ARG UNSIGNED_KMOD_2
# ARG ...

COPY --from=signimage $SIGNROOT/$UNSIGNED_KMOD_0 $UNSIGNED_KMOD_0
COPY --from=signimage $SIGNROOT/$UNSIGNED_KMOD_1 $UNSIGNED_KMOD_1
