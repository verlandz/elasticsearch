package generate

import (
	"github.com/verlandz/elasticsearch/scripts/utils"
)

// Data Source:
// https://www.michaelhartzell.com/blog/bid/17550/howto-6-list-of-strengths-talents-you-may-have

var listAbility = []string{
	"Public Speaking",
	"Writing",
	"Self Management",
	"Networking (person to person)",
	"Networking (in the virtual world)",
	"Critical Thinking",
	"Decision Making",
	"Math",
	"Research",
	"Relaxation",
	"Accounting",
	"Bookkeeping",
	"Taxes",
	"Legal",
	"Marketing",
	"Guerrilla Marketing",
	"Advertising",
	"Graphics",
	"Music",
	"Art",
	"Drawing",
	"Photography",
	"Woodworking",
	"Video Creation",
	"Website",
	"Wisdom",
	"Programming",
	"Jokes / Humor",
	"Articulate",
	"Creativity",
	"Innovation",
	"Trouble-shooter",
	"Foreign Language",
	"Academics",
	"Money Management",
	"Sign Language",
	"Enthusiasm",
	"Teaching  / Training",
	"High Energy",
	"Negotiating Skills",
	"Identify Strengths and Weaknesses",
	"Planning",
	"Leadership",
	"Listening",
	"Reading",
	"Persuasive",
	"Ability to deal with Failure",
	"Typing",
	"Initiative",
	"Strategic Planning",
	"Intuition",
	"Imaginative",
	"Future Thinking",
	"Analyzing the Past",
	"Personal Productivity",
	"Time Management ",
	"Financial Planning",
	"Ability to spot new Trends",
	"Raise Money",
	"Inspiring",
	"Story Telling",
	"Ability to make Friends",
	"Communication Skills",
	"Project Management",
	"SEO",
	"Sales",
	"Problem Solving",
	"Computer Literacy",
	"Detail Orientation",
	"Social Networking",
	"Logistics",
	"Social Intelligence",
	"Relieve Stress",
	"Asking Questions",
	"Volunteering",
	"Risk Management",
	"Hiring / Recruiting",
	"Self Control",
	"Health / Fitness",
	"Reliability",
	"Dexterity",
	"Juggling",
	"Magic",
	"Singing",
	"Visualization",
	"Adaptability",
	"Inventiveness",
	"Imagination",
	"Athleticism",
	"People Judgment",
	"Awareness",
	"Integrity / Honesty",
	"Empathy",
	"Self-Discipline",
	"Encouraging",
	"Software",
	"Computers / IT",
	"Affiliate systems",
	"Financial Management",
	"Human Resources",
	"Meeting management",
	"Learner",
	"Polyglot (learn/know a many languages)",
	"Systems management",
	"Brainstorming",
	"Positiveness",
	"Making Connections",
	"Ability to focus",
	"Ability to handle Change",
	"Conflict Resolution",
	"Self-Assurance",
	"Maintenance / Routine Tasks",
	"Futuristic",
	"Fairness",
}

const MIN_ABILITY, MAX_ABILITY = 1, 5

var lenListAbility int

func init() {
	lenListAbility = len(listAbility)
}

func GenAbility() []string {
	visited := map[string]bool{}
	result := []string{}
	n_ability := utils.RandInt(MIN_ABILITY, MAX_ABILITY)

	for {
		temp := listAbility[utils.RandInt(0, lenListAbility-1)]
		if !visited[temp] {
			visited[temp] = true
			result = append(result, temp)
			if len(result) >= n_ability {
				break
			}
		}
	}

	return result
}
