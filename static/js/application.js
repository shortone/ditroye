(function ($) {
	Friend = Backbone.Model.extend({
		name = null
	});
	
	Friends = Backbone.Collection.extend({
		initialize: function (models, options) {
			this.bind("add", options.view.addFriendLi);
		}
	});
	AppView = Backbone.View.extend({
		el: $("body"),
		events: {
			"click #add-friend": "showPrompt",
		},
		showPrompt: function () {
			var friend_name = prompt("Who is your friend?");
			var friend_model = new Friend({
				name: friend_name
			});
		}
		addFriendLi: function (model) {
			$("#friends-list").append("<li>" + model.get('name') + "</li>");
		}
	});
	var appView = new AppView;
}) (jQuery);
