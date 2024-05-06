// Returns one of the following responses:
// - https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200
// - https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408
// - https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429
// - https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/502
// - https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503
const randomResponse = (request: Request) => {
  const no = Math.random();
  const headers = new Headers();

  // Something 10 to 20 seconds in the future
  const retryAfter = new Date(
    Date.now() + Math.floor(Math.random() * 21 + 10) * 1000,
  ).toUTCString();

  if (no < 0.1) {
    return { body: ``, delay: 30_000, status: 408, headers };
  } else if (no < 0.2) {
    headers.append("Retry-After", `${retryAfter}`);
    return { body: ``, delay: 0, status: 429, headers };
  } else if (no < 0.3) {
    return { body: ``, delay: 0, status: 502, headers };
  } else if (no < 0.4) {
    headers.append("Retry-After", `${retryAfter}`);
    return { body: ``, delay: 0, status: 503, headers };
  }
  const body = `Your user-agent is: ${
    request.headers.get(
      "user-agent",
    ) ?? "Unknown"
  }`;
  return { body, delay: 0, status: 200, headers };
};

const handler = async (request: Request): Promise<Response> => {
  const { body, delay, status, headers } = randomResponse(request);

  if (delay > 0) {
    await new Promise((resolve) => setTimeout(resolve, delay));
  }

  return new Response(body, { status, headers });
};

Deno.serve({ port: 4505 }, handler);
