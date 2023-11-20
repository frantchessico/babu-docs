import axios from "axios";

const options = {
  method: 'POST',
  url: 'https://babuapi.savanapoint.com/api/core',
  headers: {
    secret_key: 'SUA_CHAVE_SECRETA_AQUI',
    'Content-Type': 'application/json'
  },
  data: {
    collectionKey: 'SUA_CHAVE_DE_COLEÇÃO_AQUI',
    fields: { 
      FIRSTNAME: 'Maria', 
      LASTNAME: 'Dapersa', 
      EMAIL: 'bamariaentosimao@gmail.com',
      // Adicione mais campos aqui conforme necessário
    }
  }
};

axios.request(options)
  .then(function (response) {
    console.log("Dados enviados com sucesso:", response.data);
  })
  .catch(function (error) {
    console.error("Erro ao enviar dados:", error);
  });