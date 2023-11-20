<?php

$url = 'https://babuapi.savanapoint.com/api/core';
$secret_key = 'SUA_CHAVE_SECRETA_AQUI';
$collectionKey = 'SUA_CHAVE_DE_COLECAO_AQUI';

$data = array(
    'collectionKey' => $collectionKey,
    'fields' => array(
        'FIRSTNAME' => 'Maria',
        'LASTNAME' => 'Dapersa',
        'EMAIL' => 'bamariaentosimao@gmail.com'
        // Adicione mais campos aqui conforme necessário
    )
);

$data_string = json_encode($data);

$headers = array(
    'Content-Type: application/json',
    'secret_key: ' . $secret_key
);

$ch = curl_init($url);

curl_setopt($ch, CURLOPT_CUSTOMREQUEST, 'POST');
curl_setopt($ch, CURLOPT_POSTFIELDS, $data_string);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);

$result = curl_exec($ch);

if (curl_errno($ch)) {
    echo 'Erro ao enviar dados: ' . curl_error($ch);
} else {
    echo 'Dados enviados com sucesso: ' . $result;
}

curl_close($ch);

?>
