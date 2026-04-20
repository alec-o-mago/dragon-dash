components {
  id: "laser"
  component: "/components/boss_bullets/laser/laser.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"laser_active\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/images/core/core.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.1
  }
  scale {
    x: 0.01
    y: 160.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"boss_bullets\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"box\"\n"
  "  }\n"
  "  data: 12.0\n"
  "  data: 650.0\n"
  "  data: 10.0\n"
  "}\n"
  "event_collision: false\n"
  "event_contact: false\n"
  "event_trigger: false\n"
  ""
}
