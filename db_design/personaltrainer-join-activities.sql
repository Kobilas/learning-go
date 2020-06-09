use PersonalTrainer;

-- Activity 1
SELECT *
FROM Exercise
JOIN ExerciseCategory
	ON Exercise.ExerciseCategoryId = ExerciseCategory.ExerciseCategoryId;

-- Activity 2
SELECT ExerciseCategory.`Name`, Exercise.`Name`
FROM Exercise
JOIN ExerciseCategory
	ON Exercise.ExerciseCategoryId = ExerciseCategory.ExerciseCategoryId
WHERE isnull(ExerciseCategory.ParentCategoryId);

-- Activity 3
SELECT
	ExerciseCategory.`Name` AS CategoryName,
    Exercise.`Name` AS ExerciseName
FROM Exercise
JOIN ExerciseCategory
	ON Exercise.ExerciseCategoryId = ExerciseCategory.ExerciseCategoryId
WHERE isnull(ExerciseCategory.ParentCategoryId);

-- Activity 4
SELECT `Client`.FirstName, `Client`.LastName, `Client`.BirthDate, Login.EmailAddress
FROM `Client`
JOIN Login
	ON `Client`.ClientId = Login.ClientId
WHERE `Client`.BirthDate BETWEEN '1990-01-01' AND '1999-12-31';

-- Activity 5
SELECT Workout.`Name`, `Client`.FirstName, `Client`.LastName
FROM `Client`
INNER JOIN ClientWorkout
	ON `Client`.ClientId = ClientWorkout.ClientId
INNER JOIN Workout
	ON ClientWorkout.WorkoutId = Workout.WorkoutId
WHERE LEFT(`Client`.LastName, 1) = "C";

-- Activity 6
SELECT
	Workout.`Name` as WorkoutName,
    Goal.`Name` as GoalName
FROM Workout
INNER JOIN WorkoutGoal
	ON Workout.WorkoutId = WorkoutGoal.WorkoutId
INNER JOIN Goal
	ON WorkoutGoal.GoalId = Goal.GoalId;