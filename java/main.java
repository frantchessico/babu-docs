import java.io.BufferedReader;
import java.io.DataOutputStream;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;

public class Main {
    public static void main(String[] args) {
        try {
            URL url = new URL("https://babuapi.savanapoint.com/api/core");
            HttpURLConnection connection = (HttpURLConnection) url.openConnection();
            connection.setRequestMethod("POST");
            connection.setRequestProperty("Content-Type", "application/json");
            connection.setRequestProperty("secret_key", "SUA_CHAVE_SECRETA_AQUI");
            connection.setDoOutput(true);

            String data = "{\"collectionKey\":\"SUA_CHAVE_DE_COLECAO_AQUI\",\"fields\":{\"FIRSTNAME\":\"Maria\",\"LASTNAME\":\"Dapersa\",\"EMAIL\":\"bamariaentosimao@gmail.com\"}}";
            // Adicione mais campos conforme necessário

            DataOutputStream outputStream = new DataOutputStream(connection.getOutputStream());
            outputStream.writeBytes(data);
            outputStream.flush();
            outputStream.close();

            int responseCode = connection.getResponseCode();
            System.out.println("Código de resposta: " + responseCode);

            BufferedReader in = new BufferedReader(new InputStreamReader(connection.getInputStream()));
            String inputLine;
            StringBuilder response = new StringBuilder();

            while ((inputLine = in.readLine()) != null) {
                response.append(inputLine);
            }
            in.close();

            System.out.println("Resposta: " + response.toString());
        } catch (Exception e) {
            System.out.println("Erro ao enviar dados: " + e);
        }
    }
}
