<?php
require('mysqli_connect.php');
//if (!empty($_POST)) {
if (empty($_POST)) {
    print ('post is empty!');
} else if (!isset($_POST)) {
    print('nothing is even getting passed in man!');
}

$name = $_POST['name'];
$course = $_POST['course'];
$grade = $_POST['grade'];
$_POST['error'] = '';
print($_POST['name']);
//method is for sanitation before you through it into the query.
$check_all = [
    $name => ['regex' => '/^[a-zA-Z0-9, ]++$/', 'method' => 'sanitized_string', 'key' => 'name'],
    $grade => ['regex' => '/^[a-zA-Z0-9, ]++$/', 'method' => 'sanitized_string', 'key' => 'course'],
    $course => ['regex' => '/^[0-9]{1,3}$/', 'method' => 'sanitized_int', 'key' => 'grade']
];

$infoadd = ['success' => true, 'data' => $_POST, 'error' => 'something went wrong'];
print(json_encode($_POST));
//}

//regex and validation here.
function regex_checker($regex, $check)
{
    return preg_match($regex, $check);
}

function string_checker($value)
{
    return addslashes($value);
}

function number_checker($value)
{
    return (int)$value;
}

function fix_stuff($value, $method)
{
    switch ($method) {
        case 'sanitized_string':
            string_checker($value);
            break;
        case 'sanitized_int':
            number_checker($value);
            break;
        default:
            return false;
            break;
    }
}

foreach ($check_all as $key => $value) {
    if (regex_checker($value['regex'], $key)) {
        $_POST[$key['key']] = fix_stuff($key, $value['method']);
    } else {
        $_POST['error'] = 'not valid regex';
    }
}

if (empty($_POST['error'])) {
//then pass the sanitized string into the database. This means the above loop should return variables, possibly into an associative array. Then throw those
// variables into where you see $name, $course, $grade below.
    $query = "INSERT INTO `students`(`name`, `course`, `grade`) VALUES ('{$_POST['name']}','{$_POST['course']}','{$_POST['grade']}')";

    mysqli_query($conn, $query);
}
?>