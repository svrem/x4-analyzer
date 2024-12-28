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

window.onload = function () {
  const anchorTags = document.querySelectorAll("a");
  anchorTags.forEach(function (a) {
    a.addEventListener("click", function (ev) {
      ev.preventDefault();
    });
  });
  const dropdownEl = document.querySelector("[data-dropdown-toggle]");
  if (dropdownEl) {
    dropdownEl.click();
  }
  const modalEl = document.querySelector("[data-modal-toggle]");
  if (modalEl) {
    modalEl.click();
  }
  const dateRangePickerEl = document.querySelector("[data-rangepicker] input");
  if (dateRangePickerEl) {
    dateRangePickerEl.focus();
  }
  const drawerEl = document.querySelector("[data-drawer-show]");
  if (drawerEl) {
    drawerEl.click();
  }
};
