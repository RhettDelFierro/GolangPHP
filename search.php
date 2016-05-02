<?php
require_once('mysqli_connect.php');
if (!empty($_POST)){
    $id = $_POST['student_id'];
    print(json_encode($_POST));
}
$query= "DELETE FROM `students` WHERE `ID` = $id LIMIT 1";

mysqli_query($conn, $query);

?>