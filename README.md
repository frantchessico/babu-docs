

# Documentação de Integração do Babu

Bem-vindo à documentação de integração do Babu. Aqui, você encontrará todas as informações necessárias para começar a usar o Babu para armazenar e gerenciar dados de seus formulários em projetos frontend.

## Passo 1: Registre-se no Babu

Antes de começar, você precisará criar uma conta no Babu. Acesse [https://babuapp.com](https://babu.savanapoint.com/) para se registrar e obter suas chaves de API.

## Passo 2: Realize a Integração

Agora que você possui suas chaves de API, pode começar a integrar seu frontend com o Babu. Utilizaremos a biblioteca Axios para fazer solicitações HTTP. Aqui está um exemplo de código para realizar uma integração básica:

```javascript
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
```

Certifique-se de substituir 'SUA_CHAVE_SECRETA_AQUI' e 'SUA_CHAVE_DE_COLEÇÃO_AQUI' pelas suas chaves de API e chave de coleção, respectivamente. Além disso, observe que o Babu suporta muitos campos, e você pode adicionar mais campos ao objeto `fields` conforme necessário, utilizando as chaves da lista `fieldKeys`.

## Campos Suportados

O Babu suporta uma ampla variedade de campos para armazenamento de dados de formulários. Você pode utilizar as seguintes chaves no objeto `fields`:

- 'FIRSTNAME'
- 'LASTNAME'
- 'FULLNAME'
- 'EMAIL'
- 'PHONE'
- 'ADDRESS'
- 'CITY'
- 'STATE'
- 'ZIPCODE'
- 'COUNTRY'
- 'BIRTHDATE'
- 'PHONE_NUMBER'
- 'GENDER'
- 'COMMENTS_OR_NOTES'
- 'COMPANY_NAME'
- 'JOB_TITLE'
- 'WEBSITE'
- 'SOCIAL_MEDIA'
- 'INTERESTS'
- 'CONTACT'
- 'LANGUAGE'
- 'IDENTIFICATION_NUMBER'
- 'IDENTIFICATION_TYPE'
- 'ATTACHMENT_FILE'
- 'USERNAME'
- 'STREET'
- 'APARTMENT_NUMBER'
- 'COUNTY'
- 'POSTAL_CODE'
- 'OCCUPATION'
- 'MARITAL_STATUS'
- 'NUMBER_OF_DEPENDENTS'
- 'INCOME'
- 'HOBBIES'
- 'MEMBERSHIP_LEVEL'
- 'SUBSCRIPTION_PLAN'
- 'REFERRAL_SOURCE'
- 'COMMENTS'
- 'PREFERRED_CONTACT_METHOD'
- 'SHIPPING_ADDRESS'
- 'BILLING_ADDRESS'
- 'ORDER_NUMBER'
- 'RETURN_REASON'
- 'SURVEY_RESPONSES'
- 'PAYMENT_METHOD'
- 'TRANSACTION_ID'
- 'ACCOUNT_NUMBER'
- 'ACCOUNT_TYPE'
- 'EMPLOYER'
- 'EMPLOYMENT_STATUS'
- 'REFERENCES'
- 'CUSTOMER_ID'
- 'SUPPORT_REQUEST_DESCRIPTION'
- 'TICKET_NUMBER'

Você pode adicionar qualquer combinação desses campos ao seu objeto `fields` ao realizar uma solicitação de integração.

## Perguntas e Suporte

Se você tiver alguma dúvida ou precisar de suporte adicional, sinta-se à vontade para entrar em contato conosco através do [suporte do Babu]([https://babuapp.com/support](https://babu.savanapoint.com/)).

Agradecemos por escolher o Babu para simplificar o armazenamento e gerenciamento de dados dos seus formulários. Esperamos que você tenha uma experiência incrível usando nossa plataforma!
