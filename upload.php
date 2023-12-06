<?php 

$table = $_FILES["table"];

//валидация




?>

<?php
if(is_dir('uploads') == false) {
    mkdir('uploads', 0777, true);
}


move_uploaded_file($table["tmp_name"],"uploads" . $table["name"]);
?>


<a href = "cfu.xlsx" download=""> Скачать таблицу</a>