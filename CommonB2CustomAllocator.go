package box2d

type B2CustomAllocator struct {
	M_bodies     []*B2Body
	M_contacts   []B2ContactInterface // has to be backed by pointers
	M_joints     []B2JointInterface   // has to be backed by pointers
	M_velocities []B2Velocity
	M_positions  []B2Position

	M_bodyCount       int
	M_contactCount    int
	M_jointCount      int
	M_velocitiesCount int
	M_positionsCount  int
}

func MakeB2CustomAllocator() B2CustomAllocator {
	return B2CustomAllocator{}
}

func (allocator *B2CustomAllocator) AllocateBodies(count int) *[]*B2Body {
	if count <= allocator.M_bodyCount {
		return &allocator.M_bodies
	}

	allocator.M_bodies = make([]*B2Body, count)
	allocator.M_bodyCount = count

	return &allocator.M_bodies
}

func (allocator *B2CustomAllocator) AllocateContacts(count int) *[]B2ContactInterface {
	if count <= allocator.M_contactCount {
		return &allocator.M_contacts
	}

	allocator.M_contacts = make([]B2ContactInterface, count)
	allocator.M_contactCount = count

	return &allocator.M_contacts
}

func (allocator *B2CustomAllocator) AllocateJoints(count int) *[]B2JointInterface {
	if count <= allocator.M_jointCount {
		return &allocator.M_joints
	}

	allocator.M_joints = make([]B2JointInterface, count)
	allocator.M_jointCount = count

	return &allocator.M_joints
}

func (allocator *B2CustomAllocator) AllocateVelocities(count int) []B2Velocity {
	if count <= allocator.M_velocitiesCount {
		return allocator.M_velocities
	}

	allocator.M_velocities = make([]B2Velocity, count)
	allocator.M_velocitiesCount = count

	return allocator.M_velocities
}

func (allocator *B2CustomAllocator) AllocatePositions(count int) []B2Position {
	if count <= allocator.M_positionsCount {
		return allocator.M_positions
	}

	allocator.M_positions = make([]B2Position, count)
	allocator.M_positionsCount = count

	return allocator.M_positions
}
