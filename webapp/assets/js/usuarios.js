$('button[id*=parar-de-seguir]').each((index, btn) => {
    $(btn).on('click', pararDeSeguir)
})
$('button[id*=seguir-]:not([id*=parar-de-seguir])').each((index, btn) => {
    $(btn).on('click', seguir)
})

function pararDeSeguir(){
    const btn = $(this)
    const usuarioId = btn[0].dataset.usuarioId
    btn.prop('disabled', true)

    $.ajax({
        url: `/usuarios/${usuarioId}/parar-de-seguir`,
        method: "POST"
    }).done(function(){
        window.location.reload()
    }).fail(function(){
        Swal.fire(
            'Ops...',
            'Erro ao parar de seguir usuário!',
            'error'
        )
        btn.prop('disabled', false)
    })
}

function seguir(e){
    const btn = $(this)
    const usuarioId = btn[0].dataset.usuarioId
    btn.prop('disabled', true)

    $.ajax({
        url: `/usuarios/${usuarioId}/seguir`,
        method: "POST"
    }).done(function(){
        window.location.reload()
    }).fail(function(){
        Swal.fire(
            'Ops...',
            'Erro ao seguir usuário!',
            'error'
        )
        btn.prop('disabled', false)
    })
}