#!/bin/sh

TEMPLATE_DIR=${PROJECT_DIR}/template/latex
TMP_DIR=`mktemp -d`

clean()
{
    rm -rf "${TMP_DIR}"
}

trap clean EXIT
cp "${TEMPLATE_DIR}/beamerposter.sty" "${TMP_DIR}"
cp "${TEMPLATE_DIR}/beamerthemeruhuisstijlposter.cls" "${TMP_DIR}"
cp "${TEMPLATE_DIR}/logo.png" "${TMP_DIR}"
mv "/tmp/${1}.tex" "${TMP_DIR}"

cd "${TMP_DIR}"
pdflatex -interaction=batchmode "${1}.tex"
mv "${1}.pdf" /tmp
