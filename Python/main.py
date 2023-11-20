from fastapi import FastAPI, HTTPException
from pydantic import BaseModel

import httpx

app = FastAPI()

class UserData(BaseModel):
    FIRSTNAME: str
    LASTNAME: str
    EMAIL: str

@app.post("/send-data")
async def enviar_dados(user_data: UserData):
    url = 'https://babuapi.savanapoint.com/api/core'
    collection_key = 'SUA_CHAVE_DE_COLECAO_AQUI'
    secret_key = 'SUA_CHAVE_SECRETA_AQUI'

    data = {
        'collectionKey': collection_key,
        'fields': {
            'FIRSTNAME': user_data.FIRSTNAME,
            'LASTNAME': user_data.LASTNAME,
            'EMAIL': user_data.EMAIL,
            # Adicione mais campos conforme necessário
        }
    }

    headers = {
        'Content-Type': 'application/json',
        'secret_key': secret_key
    }

    try:
        async with httpx.AsyncClient() as client:
            response = await client.post(url, json=data, headers=headers)
            if response.status_code == 201:
                return {"message": "Dados enviados com sucesso", "response": response.json()}
            else:
                detail = {
                    "message": "Erro ao enviar dados",
                    "response_status_code": response.status_code,
                    "response_content": response.content.decode("utf-8")
                }
                raise HTTPException(status_code=response.status_code, detail=detail)
    except httpx.HTTPError as e:
        raise HTTPException(status_code=500, detail=f"Erro na solicitação: {e}")
