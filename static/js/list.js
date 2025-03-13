document.addEventListener('DOMContentLoaded',() => {
  (function(win, doc, $){
    "using strict";
    document.querySelectorAll("td[data-id='Reply']").forEach(element => {
      element.removeAttribute('onclick')
  
      let id = element.parentElement.querySelector('.search.pointer a').getAttribute('data-id')

      element.querySelector('a').addEventListener('click', () => {
        // Swal.fire({
        //   title: "Replying to Message",
        //   html: `
        //     <h3 style="text-align: start;" class="display-5 font-weight-normal">Replying To: <span class="fw-bold">Bold text.</span></h3>
        //     <textarea rows="10" class="h-auto form-control" style="resize: none;" name="" id="" ></textarea>
        //   `,

        //   width: '800px', 
        // });

        Swal.fire({
          title: "<h3>Replying to Message</h3>",
          html: `
            <p style="text-align: start; font-size: 1.5rem" class="font-weight-normal">Replying To: <span class="fw-bold">Bold text.</span></p>
            <textarea rows="10" class="h-auto form-control" style="resize: none;" name="" id="" ></textarea>
          `,
          width: '500px',
          showCloseButton: true,
          showCancelButton: true,
          confirmButtonText: `
            <div style="margin: 0; padding: 0;">Send Message<span style="padding-left:0.5rem" class="glyphicon glyphicon-send" aria-hidden="true"></span></div>
          `,
          cancelButtonText: `
            <p style="margin:0">qweqwe</p>
          `,
        });
      })
    })
  
  })(window, document, $);
})