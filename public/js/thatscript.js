var thatApp = angular.module("thatApp", []);

thatApp.directive("dumbPassword", function () {
  var validElement = angular.element('<div>{{model.input}}</div>');

  var link = function (scope, element) {
    scope.$watch("model.input", function (value) {
      if(value === "password") {
        console.log(element);
        validElement.addClass("btn");
      } else {
        validElement.removeClass("btn");
			}
    });
  };

  return {
    restrict: "E",
    replace: true,
    template: '<div><input type="text" ng-model="model.input">',
    compile: function (tElem) {
      tElem.append(validElement);
        
      return link;
    }  
  }; 
});

