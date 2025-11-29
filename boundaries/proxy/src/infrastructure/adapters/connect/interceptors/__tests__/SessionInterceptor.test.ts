import { describe, it, expect } from "vitest";
import { createSessionInterceptor } from "../SessionInterceptor.js";

describe("createSessionInterceptor", () => {
  it("sets x-user-id header on outgoing request", async () => {
    const interceptor = createSessionInterceptor("proxy-user");
    const headerMap = new Map<string, string>();
    const mockHeader = {
      get: (key: string): string | undefined => headerMap.get(key),
      set: (key: string, value: string): void => {
        headerMap.set(key, value);
      },
    };

    const request = {
      header: mockHeader,
    } as any;

    const next = async (req: any) => {
      return req;
    };

    await interceptor(next)(request);

    expect(request.header.get("x-user-id")).toBe("proxy-user");
  });

  it("throws when service user id is empty", () => {
    expect(() => createSessionInterceptor("")).toThrow(
      /SERVICE_USER_ID is required/
    );
  });
});
