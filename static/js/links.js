const updateLinks = () => {
  const links = document.querySelectorAll("a");
  links.forEach((link) => {
    href = link.getAttribute("href");
    if (href.startsWith("/")) {
      link.setAttribute("hx-get", href + "?c=true");
      link.setAttribute("hx-push-url", href);
      link.setAttribute("hx-swap", "outerHTML");
      link.setAttribute("hx-target", "main");
    }
  });
};

const observer = new MutationObserver((mutations) => {
  mutations.forEach((mutation) => {
    if (mutation.type === "childList") {
      updateLinks();
    }
  });
});

observer.observe(document.body, {
  childList: true,
  subtree: true,
});

updateLinks();
