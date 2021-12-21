$('#signup-form').on('submit', createUser);

function createUser(event){
    event.preventDefault();

    if ($('#password').val() != $('#confirm-password').val()){
        alert('Passwords do not match');
        console.log('createUser');
        return;
    }

    $.ajax({
        url: '/users',
        method: 'POST',
        data: {
            name: $('#name').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            password: $('#password').val()
        },
	
	})
}

