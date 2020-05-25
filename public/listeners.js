window.addEventListener(`load`, ()=>{
    console.log('listeners started');
    var dashLinks = $('a.dash-link');
    dashLinks.on('click', (e)=>{
        if(e.target.id == 'home')
        {
            home()
        }
        else if(e.target.id == 'account-options')
        {
            console.log('->account-options')
        }
        else if(e.target.id == 'find-items')
        {
            console.log('->find-items')
        }
        else if(e.target.id == 'view-posted')
        {
            console.log('->view-posted')
        }
    })
    console.log('listeners ended');
})