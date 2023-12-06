<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Форма регистрации</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css">
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <div class = "container mt-4" >
        <?php
        if($_COOKIE['user'] == ''):
        ?>
        <div class="row">

            
            </div>
            <div class="col">
                <h1>Форма авторизации</h1>
        <form action="auth.php" method = "post">
            <input type = "text" class = "form-control" name = "login" id = "login" placeholder="Введите логин"><br>
            <input type = "text" class = "form-control" name = "password" id = "password" placeholder="Введите пароль"><br>
            <button class = "btn btn-success" type = "submit">Авторизоваться</button>
        </form>
            <p> Если ты только пользователь, то есть ученик и не админ, то вот тебе ссылка на твою <a href = "https://vk.com/vernadskycfu">страницу</a></p>
            </div>
            
        
            <?php else :
                ?>
                <p>Привет  <?=$_COOKIE['user']?>. Чтобы выйти нажмите здесь <a href = "/exit.php">здесь</a>.Также, если ты хочешь можешь преобразовать свою таблицу <a href = "third.php">здесь</a>.</p>
            <?php endif;?>
            
        </div>
    </div>
    
</body>
</html>