<?php
require_once('mysqli_connect.php');

$query = "SELECT * FROM `students`";

$rows = mysqli_query($conn, $query);

if (mysqli_num_rows($rows) > 0){
    while($row = mysqli_fetch_assoc($rows)){
        $output[] = $row;
    }
}
$returnedbyjson = ['success' => true, 'data' => $output, 'error' => 'something went wrong'];

print(json_encode($returnedbyjson));
?>