package http;

import retrofit2.Call;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;
import retrofit2.http.Body;
import retrofit2.http.GET;
import retrofit2.http.POST;

import java.io.IOException;
import java.util.List;

public class LinkApp {
  private final API api;

  public interface API {
    @GET("links")
    Call<List<Link>> list();

    @POST("link")
    Call<Link> add(@Body Link data);
  }

  public LinkApp() {
    String BASE_URL = "http://localhost:7070/api/";

    // Create a very simple REST adapter which points the API.
    Retrofit retrofit = new Retrofit.Builder()
      .baseUrl(BASE_URL)
      .addConverterFactory(GsonConverterFactory.create())
      .build();

    // Create an instance of our API interface.
    this.api = retrofit.create(API.class);
  }

  public Link AddLink(String url) throws IOException {
    Link payload = new Link(url, "");
    Call<Link> resp = this.api.add(payload);

    return resp.execute().body();
  }
}
