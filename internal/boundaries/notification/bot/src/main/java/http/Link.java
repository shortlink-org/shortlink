package http;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class Link {
  @SerializedName("url")
  @Expose
  private String url;

  @SerializedName("hash")
  @Expose
  private String hash;

  public String getUrk() {
    return url;
  }

  public void setUrl(String url) {
    this.url = url;
  }

  public String getHash() {
    return hash;
  }

  public void setHash(String hash) {
    this.hash = hash;
  }

  public Link(String url, String hash) {
    this.url = url;
    this.hash = hash;
  }
}
