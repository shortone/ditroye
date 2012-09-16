(function ($) {

	Friend = Backbone.Model.extend({
		//Create a model to hold friend atribute
		name: null,
		age: null
	});

	Friends = Backbone.Collection.extend({
		//This is our Friends collection and holds our Friend models
		initialize: function (models, options) {
			this.bind("add", options.view.addFriendLi);
			//Listen for new additions to the collection and call a view function if so
		}
	});

	SearchQuery = Backbone.Model.extend({
		battleTag: null
	});

	Hero = Backbone.Model.extend({
		name: null,
		level: null,
		hardcore: null,
		cl: null,
		gender: null
	});

	Heroes = Backbone.Collection.extend({
		initialize: function(models, options) {
			this.bind("add", options.view.addHero);
		}
	});
	
	PlayerProfile = Backbone.Model.extend({
		id: null,
		battleTag: null,
		urlRoot: 'http://localhost:3001/resources/profile/'
	});

	AppView = Backbone.View.extend({
		el: $("body"),
		initialize: function () {
			this.friends = new Friends( null, { view: this });
			//Create a friends collection when the view is initialized.
			//Pass it a reference to this view to create a connection between the two
		},
		events: {
			"click #search-profile": "searchProfile"
		},
		showPrompt: function () {
			var friend_name = prompt("Who is your friend?");
			var friend_model = new Friend({ name: friend_name });
			//Add a new friend model to our friend collection
			this.friends.add( friend_model );
		},
		searchProfile: function () {
			var battleTag = $("#battle-tag").val();
			if (battleTag != '') {
				console.log("Search launch for battle-tag: " + battleTag);
				var playerProfile = new PlayerProfile({
					id: battleTag,
					battleTag: battleTag
				});
				playerProfile.fetch({
					success: function (data) {
						console.log(data);
					}
				});
			} else {
				console.log("Empty battle-tag provided");
			}
		},
		addFriendLi: function (model) {
			//The parameter passed is a reference to the model that was added
			$("#friends-list").append("<li>" + model.get('name') + "</li>");
			//Use .get to receive attributes of the model
		},
		addHero: function (model) {
			console.log("new hero added: " + model.get('name'));
		}
	});

	var appview = new AppView;
})(jQuery);