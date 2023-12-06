<form method = "post" action = "upload.php" enctype = "multipart/form-data">
    <input type = "file" name = "table">
    <button type = "submit">upload</button>
</form>

<p> Ваш файл должен быть одноименным! Например: он должен быть cfu.xlsx(только такое название, для удобства)</p>
<style> 
    * {
        font-size: 26px;
    }

    input, button {
        display: block;
        margin-bottom: 10px;
    }
</style>