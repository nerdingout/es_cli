package templates

const Test = `import %s from ".";

describe("%s", () => {
  test("tests function", () => {
    expect(%s("hello")).toBe("hello");
  });
});
`
