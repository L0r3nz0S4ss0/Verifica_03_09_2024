import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.PreparedStatement;
import java.sql.SQLException;
import java.sql.Timestamp;
import java.util.ArrayList;
import java.util.List;

import org.json.JSONArray;
import org.json.JSONObject;

public class Main {
    public static void main(String[] args) {
        InputStreamReader inp = new InputStreamReader(System.in);
        BufferedReader bf = new BufferedReader(inp);
        Connection conn = null;
        try {
            conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/esercizio", "root", "");

            List<JSONObject> allComments = new ArrayList<>();
            System.out.println("Selezionare il numero di comments da prelevare");
            int nGet = Integer.parseInt(bf.readLine());


            for (int postId = 1; postId <= nGet; postId++) {
                String url = "https://jsonplaceholder.typicode.com/comments?id=" + postId;
                URL obj = new URL(url);
                HttpURLConnection con = (HttpURLConnection) obj.openConnection();
                con.setRequestMethod("GET");

                BufferedReader in = new BufferedReader(new InputStreamReader(con.getInputStream()));
                StringBuilder response = new StringBuilder();
                String inputLine;
                while ((inputLine = in.readLine()) != null) {
                    response.append(inputLine);
                }
                in.close();

                JSONArray jsonArray = new JSONArray(response.toString());

                for (int i = 0; i < jsonArray.length(); i++) {
                    JSONObject comment = jsonArray.getJSONObject(i);
                    allComments.add(comment);
                }
            }

            String insertQuery = "INSERT INTO comments (postId, id, name, email, body, data_ora) VALUES (?, ?, ?, ?, ?, ?)";
            PreparedStatement pstmt = conn.prepareStatement(insertQuery);

            for (JSONObject comment : allComments) {
                pstmt.setInt(1, comment.getInt("postId"));
                pstmt.setInt(2, comment.getInt("id"));
                pstmt.setString(3, comment.getString("name"));
                pstmt.setString(4, comment.getString("email"));
                pstmt.setString(5, comment.getString("body"));
                pstmt.setTimestamp(6, new Timestamp(System.currentTimeMillis()));

                pstmt.executeUpdate();
            }

            pstmt.close();

            System.out.println("Dati memorizzati nel database con successo.");

        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            // Chiude la connessione
            if (conn != null) {
                try {
                    conn.close();
                } catch (SQLException e) {
                    e.printStackTrace();
                }
            }
        }
    }
}
