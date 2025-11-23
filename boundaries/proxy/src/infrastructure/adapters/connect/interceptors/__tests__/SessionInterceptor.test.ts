import { describe, it, expect } from "vitest";
import { createSessionInterceptor } from "../SessionInterceptor.js";

describe("createSessionInterceptor", () => {
  it("sets user-id header on outgoing request", async () => {
    const interceptor = createSessionInterceptor("proxy-user");
    const request = {
      header: new Headers(),
    } as any;

    const next = async (req: any) => {
      return req;
    };

    await interceptor(next)(request);

    expect(request.header.get("user-id")).toBe("proxy-user");
  });

  it("throws when service user id is empty", () => {
    expect(() => createSessionInterceptor("")).toThrow(
      /SERVICE_USER_ID is required/
    );
  });
});
