function selectImage() {
    document.getElementById('fileInput').click();

    // when file imported
    document.getElementById('fileInput').onchange = function () {
        document.getElementById('imageForm').submit();
    }
}

