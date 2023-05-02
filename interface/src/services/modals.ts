export function confirmModal(message: string, callback: (value: boolean) => void) {
    const html = /*html*/ `
      <div tabIndex="-1" class="z-50 min-h-screen w-screen inset-0 bg-gray-700 fixed bg-opacity-10 flex flex-row justify-center items-center">
      <div class="relative p-4 w-full max-w-lg  ">
          <div class="relative bg-white rounded-lg shadow ">
              <p class="px-5 pt-4  text-xl">${message}</p>
              <div class="p-10 text-center flex flex-row justify-end space-x-2">
                <button class="text-white bg-azul-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center btn-yes">Sim</button>
                <button class="text-gray-500 bg-white hover:bg-gray-100 focus:ring-4 focus:outline-none focus:ring-blue-300 rounded-lg border border-gray-200 text-sm font-medium px-5 py-2.5 hover:text-gray-900 focus:z-10 btn-no">Não</button>
              </div>
          </div>
      </div>
    </div>
    `
    const modal = document.createElement("div");
    modal.innerHTML = html;
    modal.querySelector(".btn-yes")?.addEventListener("click", function () {
        modal.remove();
        callback(true);
    });
    modal.querySelector(".btn-no")?.addEventListener("click", function () {
        modal.remove();
        callback(false);
    });
    document.body.appendChild(modal);
}

type ModalPromptParams = {
    message: string,
    placeholder?: string
}

export function promptModal({ message, placeholder }: ModalPromptParams, callback: (value: string) => void) {
    const html = /*html*/ `
      <div tabIndex="-1" class="z-50 min-h-screen w-screen inset-0 bg-gray-700 fixed bg-opacity-10 flex flex-row justify-center items-center">
          <div class="relative p-4 w-full max-w-lg  ">
              <div class="relative bg-white rounded-xl shadow ">
                  <div class="text-right w-full px-3 pt-2">
                    <i class="fa-solid fa-xmark btn-close text-xl cursor-pointer"></i>
                  </div>
               
                
                  <div class="px-5 pb-6">
                      <p class="mb-2">${message}</p>
                      <input class="w-full hadow appearance-none border text-gray-700 focus:outline-none focus:shadow-outline border-genial2 rounded-full px-3 py-2 mb-2 ipt-value" placeholder="${placeholder}"/>
                      <button class="bg-genial2 hover:bg-teal-700 text-white font-bold py-2 w-full px-4 rounded-full btn-ok">OK</button>
                  </div>
              </div>
          </div>
      </div>
    `

    const modal = document.createElement("div");
    modal.innerHTML = html;
    const ipt = modal.querySelector(".ipt-value") as HTMLInputElement
    modal.querySelector(".btn-close")?.addEventListener("click", function () {
        modal.remove();
        callback("");
    });
    modal.querySelector(".btn-ok")?.addEventListener("click", function () {
        modal.remove();
        const text = ipt?.value
        callback(text)
    });
    ipt.addEventListener("keydown", function (e) {
        if (e.key === 'Enter') {
            modal.remove();
            const text = ipt.value
            callback(text)
        }
    })
    document.body.appendChild(modal);
    ipt.focus()
}


export function imageModal(src: string) {
    const html = /*html*/ `
        <div id="modal"
            class="fixed top-0 left-0 w-screen h-screen bg-black/70 flex justify-center items-center" style="z-index: 999999;">
            <a class="fixed z-90 top-6 right-8 text-white text-5xl font-bold link-close-img-modal" href="javascript:void(0)">&times;</a>
            <img id="modal-img" src="${src}" class="max-w-[800px] max-h-[600px] object-cover" />
        </div>
    `
    const modal = document.createElement("div");
    modal.innerHTML = html;
    modal.querySelector(".link-close-img-modal")?.addEventListener("click", function () {
        modal.remove();
    });
    document.body.appendChild(modal);
}

