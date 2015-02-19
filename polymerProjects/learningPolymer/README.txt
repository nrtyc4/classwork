Table of contents
    The <template> element
    Templates with data binding
    Dynamic, two-way data binding
  k  Event handling and data binding

*For Polymer elements, the model is always the element itself. For example, consider the  simpleElement.html example.
    In the example, the owner property is the model for the name-tag element.
    If you update the owner property, you change the contents of the tag.

*In the data binding example, the template can be broken down as follows:
    The repeat="{{s in salutations}}" tells the template to generate a DOM fragment (or instance) for each element in the salutations array.
    The contents of the template define what each instance looks like. In this case, it contains a <li> with a text node and an <input> as its children.
    The expressions {{s.what}} and {{s.who}} create data bindings to objects in the salutations array.

