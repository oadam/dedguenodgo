$(document).ready(function() {

	var loginObservable = createStorageObservable(sessionStorage, 'login');
	var welcomeHidden = createStorageObservable(sessionStorage, 'welcomeHidden');

	window.server = new Server({
		invalidLoginCallback: function() {
			loginObservable(null);
		}
	});
	//little hack to wake up appengine
	server.wakeUp();
	window.appViewModel = new AppViewModel({
		server: window.server,
		logoutCallback: function() {
			loginObservable(null);
		},
		confirm: function(t) {
			return confirm(t);
		},
		prompt: function(t) {
			return prompt(t);
		}
	});
	window.loginViewModel = new LoginViewModel({
		server: server,
		loginCallback: function(login) {
			loginObservable(login);
		}
	});
	var appDiv = $('#app');
	var welcomeDiv = $('#welcome');
	var loginDiv = $('#login');
	var loadingDiv = $('#loading'); 
	var errorDiv = $('#error'); 

	ko.applyBindings(window.appViewModel, appDiv[0]);
	ko.applyBindings(window.loginViewModel, loginDiv[0]);
	ko.applyBindings({
		welcomeHidden: welcomeHidden,
		hideWelcome: function() {
			welcomeHidden(true);
		}
	}, welcomeDiv[0]);

	function onLogin(login) {
		server.setLogin(login);

		if (!login || !login.userId) {
			appDiv.hide();
			loginDiv.show();
			loadingDiv.hide();
			errorDiv.hide();
		} else {
			appDiv.hide();
			loginDiv.hide();
			loadingDiv.show();
			errorDiv.hide();


			server.getUsersAndPresents().always(function() {
				loadingDiv.hide();
			}).fail(function(e) {
				errorDiv.show();
				console.error(e);
			}).done(function(usersAndPresents) {
				appDiv.show();
				appViewModel.users(usersAndPresents.users);
				appViewModel.presents(usersAndPresents.presents);
				appViewModel.loggedInUser(login.userId);
			});
		}
	}
	loginObservable.subscribe(onLogin);
	onLogin(loginObservable());

	//reorderability
	var list = document.getElementById('present-list');
	new Slip(list);
	$(list)
	//prevent swiping
	.on('slip:beforeswipe', function(e) {
		e.preventDefault();
	})
	//without this, reordering requires to hold mouse down for 1 second
	.on('slip:beforewait', function(e) {
		e.preventDefault();
	})
	.on('slip:beforereorder', function(e) {
		if ($(e.target).is('button, a')) {
			e.preventDefault();
		}
	})
	.on('slip:reorder', function(e) {
		var d = e.originalEvent.detail;
		var from = d.originalIndex;
		var to = d.spliceIndex;
		appViewModel.reorder(from, to);
	});
});
